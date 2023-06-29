<template>
  <BasicTable @register="registerTable">
  </BasicTable>
</template>

<script setup lang="ts">
import { BasicColumn, BasicTable, useTable } from "/@/components/Table";
import { adminList } from "@/api/custom/admin";
import { formatToDateTime } from "@/utils/dateUtil";


const columns: BasicColumn[] = [
  {
    title: "账号",
    dataIndex: "userName",
  },
  {
    title: "密码",
    dataIndex: "password",
  },
  {
    title: "盐",
    dataIndex: "salt",
  },
  {
    title: "昵称",
    dataIndex: "nickName",
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
    labelWidth: 50,
    schemas: [
      {
        field: `userName`,
        label: `账号`,
        component: 'Input',
      },
      {
        field: `nickName`,
        label: `昵称`,
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
