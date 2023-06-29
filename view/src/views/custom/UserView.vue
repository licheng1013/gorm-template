<template>
  <BasicTable @register="registerTable">
  </BasicTable>
</template>

<script setup lang="ts">
import { BasicColumn, BasicTable, ColumnChangeParam, useTable } from "/@/components/Table";
import { adminList } from "@/api/custom/admin";
import { formatToDate, formatToDateTime } from "@/utils/dateUtil";


const columns: BasicColumn[] = [
  {
    title: "账号",
    dataIndex: "tel",
  },
  {
    title: "昵称",
    dataIndex: "nickname",
  },
  {
    title: "头像",
    dataIndex: "avatar",
  },
  {
    title: "性别",
    dataIndex: "sex",
    format: (text, record) => {
      if (record.sex === 1) {
        return "男"
      }else if (record.sex === 2) {
        return "女"
      }
      return "未知"
    }
  },
  {
    title: "创建时间",
    dataIndex: "createTime",
    format: text => formatToDateTime(text)
  }
];


const [registerTable, { reload }] = useTable({
  canResize:true,
  loading:false,
  striped:false,
  bordered:false,
  showTableSetting:true,
  useSearchForm:true,
  tableSetting: {
    setting:false,
    fullScreen: true
  },
  formConfig: {
    schemas: [
      {
        field: `tel`,
        label: `账号`,
        component: 'Input',
      }
    ],
  },
  title: '用户列表',
  titleHelpMessage:"温馨提醒",
  api: async (params) => {
    params.size = params.pageSize
    console.log(params);
    let data = await adminList(params)
    return {items:data.list, total: data.total}
  },
  columns: columns,
  pagination: { pageSize: 10},
});
</script>

<style scoped lang="less">

</style>
