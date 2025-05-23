<!-- eslint-disable no-continue -->
<script setup lang="tsx">
import type { Ref } from 'vue';
import { computed, h, ref } from 'vue';
import { NSpace, NTag } from 'naive-ui';
import dayjs from 'dayjs';
import { deleteDnatPolicy, getDnatPolicyList } from '@/service/api';
import { $t } from '@/locales';
import CustomTableHeader from '@/components/table/custom-table-header.vue';
import { ADD_FIREWALL_POLICY } from '@/utils/permissions_consts';
import { checkIpAddr, checkIpInNet } from '@/utils/ip_check';
import { createAction } from '@/components/policy/utils';
import UpdatePosition from './modules/update-position.vue';
import DataSearch from './modules/table-search.vue';
import AddData from './modules/add-data.vue';
import UpdateData from './modules/update-data.vue';

interface SearchParams {
  network: string;
  tip: string;
  destIp: string;
}

const listParams: Ref<SearchParams> = ref({
  network: '',
  tip: '',
  destIp: ''
});

function resetSearchParams() {
  listParams.value = {
    network: '',
    tip: '',
    destIp: ''
  };
  getData();
}
let data: any[] = [];
const loading = ref(false);
const showUpdateData = ref(false);
const updateData = ref<any>({});
const showUpdateDataPosition = ref(false);

const network = ref<any>([]);

function updateModalClick(row: any, method: string) {
  if (method === 'update') {
    updateData.value = row;
    showUpdateData.value = true;
  } else if (method === 'delete') {
    window.$dialog?.warning({
      title: $t('common.dialog.warning'),
      content: () =>
        h('p', null, [
          h('span', null, { default: () => $t('common.confirmDelete') })
          // h('p', { style: 'color: red;margin-top:10px' }, { default: () => row.loginname })
        ]),
      positiveText: $t('common.confirm'),
      negativeText: $t('common.cancel'),
      onPositiveClick: async () => {
        const { error } = await deleteDnatPolicy(row.id);
        if (error) return;
        getData();
        window.$message?.success($t('common.deleteSuccess'));
      }
    });
  } else if (method === 'position') {
    updateData.value = row;
    showUpdateDataPosition.value = true;
  }
}

const columns = computed<any>(() => [
  {
    key: 'id',
    title: 'id',
    align: 'center',
    show: true,
    width: 64
  },

  {
    show: true,
    key: 'comment',
    title: $t('datatable.description'),
    align: 'center',
    minWidth: 200
  },

  {
    show: true,
    title: $t('page.firewallPolicy.sourceEthernet'),
    align: 'center',
    minWidth: 100,
    key: 'iif'
  },

  {
    show: true,
    title: $t('page.firewallPolicy.destIp'),
    align: 'center',
    minWidth: 200,
    render(row: any) {
      if (row.dip && row.dip !== '') {
        return h(
          NSpace,
          { justify: 'center' },
          {
            default: () =>
              row.dip.split(',').map((x: string) => {
                return h(
                  NTag,
                  {
                    bordered: false,
                    type: 'success'
                  },
                  {
                    default: () => x
                  }
                );
              })
          }
        );
      }

      return h(
        NTag,
        {
          bordered: false,
          type: 'success'
        },
        {
          default: () => $t('page.firewallPolicy.allIp')
        }
      );
    }
  },

  {
    show: true,
    title: $t('page.firewallPolicy.nat'),
    align: 'center',
    minWidth: 200,
    key: 'dnat'
  },

  {
    show: true,
    title: $t('page.firewallPolicy.port'),
    align: 'center',
    minWidth: 100,
    render(row: any) {
      return h(
        NSpace,
        { justify: 'center' },
        {
          default: () =>
            row.port.map((x: any) => {
              let protocol = x.protocol;
              if (!protocol || protocol === '') {
                protocol = row.protocol;
              }
              return h(
                NTag,

                {
                  bordered: false,
                  type: 'success'
                },
                {
                  default: () => `${protocol}: ${x.key} → ${x.value}`
                }
              );
            })
        }
      );
    }
  },

  {
    show: true,
    title: $t('datatable.createTime'),
    key: 'createdAt',
    align: 'center',
    width: 200,
    render(row: any) {
      return h('span', null, {
        // format(row.createdAt, 'yyyy-MM-dd HH:mm:ss')
        default: () =>
          row.createdAt && row.createdAt > 0 ? dayjs.unix(row.createdAt).format('YYYY-MM-DD HH:mm:ss') : ''
      });
    }
  },

  {
    show: true,
    title: $t('datatable.action'),
    key: 'actions',
    align: 'center',
    width: 200,
    render: (row: any) => createAction(row, updateModalClick)
  }
]);

const pagination = ref<any>({
  page: 1,
  pageSize: 20,
  showSizePicker: true,
  displayOrder: ['size-picker', 'pages', 'quick-jumper'],
  itemCount: 0,
  showQuickJumper: true,
  pageSizes: [10, 20, 50, 100],
  onChange: (page: number) => {
    pagination.value.page = page;
  },
  onUpdatePageSize: (pageSize: number) => {
    pagination.value.pageSize = pageSize;
    pagination.value.page = 1;
  },
  prefix({ itemCount }: { itemCount: number }) {
    return $t('datatable.itemCount', { total: itemCount });
  }
});
const filterData = ref<any>([]);
const tableRef = ref<any>(null);

async function getData() {
  loading.value = true;
  const { data: res, error } = await getDnatPolicyList();
  loading.value = false;
  if (error) return;
  network.value = res.network;
  if (!res.data) {
    data = [];
    pagination.value.itemCount = 0;
    filterData.value = [];
    return;
  }
  data = res.data;

  filterDataFn();
  // pagination.value.itemCount = res.total;
}

getData();

const showAddData = ref(false);
//
function handleAdd() {
  showAddData.value = true;
}

//  过滤
function filterDataFn() {
  // 翻页归1
  if (tableRef.value) tableRef.value.page(1);
  loading.value = true;
  //  只显示符合条件的数据
  // eslint-disable-next-line complexity
  filterData.value = data.filter((item: any) => {
    if (listParams.value.tip && listParams.value.tip.trim() !== '') {
      // 判断ip 是否有效，有效才进行过滤
      if (checkIpAddr(listParams.value.tip)) {
        if (item.dnat !== listParams.value.tip.trim()) return false;
      }
    }
    if (item.dip && item.dip !== '' && listParams.value.destIp && listParams.value.destIp.trim() !== '') {
      // 判断ip 是否有效，有效才进行过滤
      if (checkIpAddr(listParams.value.destIp)) {
        // 判断是否在范围&& checkIpInNet()
        const tmpIps = item.dip.split(',');

        let tmpIpState = false;
        for (let i = 0; i < tmpIps.length; i += 1) {
          // eslint-disable-next-line max-depth
          if (tmpIps[i] && checkIpInNet(listParams.value.destIp, tmpIps[i])) {
            tmpIpState = true;
            break;
          }
        }
        if (!tmpIpState) return false;
      }
    }

    if (listParams.value.network && listParams.value.network.trim() !== '') {
      if (item.iif !== listParams.value.network) return false;
    }

    return true;
  });

  loading.value = false;
}
</script>

<template>
  <NSpin :show="loading">
    <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
      <DataSearch v-model:model="listParams" :network="network" @reset="resetSearchParams" @search="filterDataFn" />
      <NCard :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
        <template #header-extra>
          <CustomTableHeader
            v-model:columns="columns"
            :add-permission="ADD_FIREWALL_POLICY"
            @add="handleAdd"
            @refresh="getData"
          />
        </template>

        <template #header>
          <div class="mb-5px mt-5px">
            {{ $t('page.firewallPolicy.list') }}
          </div>
        </template>
        <NDataTable
          ref="tableRef"
          pagination-behavior-on-filter="first"
          :columns="columns"
          :data="filterData"
          :pagination="pagination"
        />
      </NCard>
    </div>

    <AddData v-model:show="showAddData" :network="network" @close="getData" />
    <UpdateData v-model:show="showUpdateData" :row="updateData" :network="network" @close="getData" />
    <UpdatePosition v-model:show="showUpdateDataPosition" :row="updateData" @close="getData" />
  </NSpin>
</template>

<style scoped lang="scss">
:deep(.n-pagination) {
  .n-pagination-prefix {
    margin-right: 10px;
  }
  .n-select {
    margin-right: 10px;
  }
  .n-pagination-quick-jumper {
    margin-left: 20px;
  }
}
</style>
