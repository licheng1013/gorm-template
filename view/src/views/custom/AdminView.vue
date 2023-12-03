<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="openModal">
          添加
        </a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                label: '删除',
                icon: 'ic:outline-delete-outline',
                onClick: handleDelete.bind(null, record),
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <BasicModal @register="registerAdd" v-bind="$attrs" title="添加" :helpMessage="['添加记录']"
                width="700px" @ok="onCloseModal">
      <BasicForm @register="registerForm" @submit="handleSubmit"/>
    </BasicModal>
  </div>

</template>

<script setup lang="ts">
import {BasicColumn, BasicTable, FormSchema, useTable,TableAction} from "/@/components/Table";
import {adminDelete, adminInsert, adminList} from "@/api/custom/admin";
import {formatToDateTime} from "@/utils/dateUtil";
import {useMessage} from "@/hooks/web/useMessage";
import { BasicModal, useModal } from "@/components/Modal";
import {BasicForm,useForm} from "@/components/Form";

const { createMessage, createConfirm } = useMessage();
const {success} = createMessage;
const [registerAdd, {openModal, closeModal}] = useModal();

const schemas: FormSchema[] = [
  {
    field: "userName",
    component: "Input",
    label: "账号",
    required: true
  },
  {
    field: "password",
    component: "Input",
    label: "密码",
    required: true
  },
  {
    field: "salt",
    component: "Input",
    label: "盐",
    required: true
  },
  {
    field: "nickName",
    component: "Input",
    label: "昵称",
    required: true
  }
];

const handleSubmit = (values: Recordable) => {
  // 时间转换
  console.log(values);
  adminInsert(values).then(res => {
    // 插入成功
    success("Success message");
    reload();
    resetFields();
  });

};

const onCloseModal = () => {
  closeModal();
  submit();
};


const [registerForm, {submit, resetFields}] = useForm({
  schemas,
  showActionButtonGroup: false,
  labelWidth: 120,
  size: "large",
  baseColProps: {
    xs: 22
  }
});


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


const [registerTable, {reload}] = useTable({
  canResize: true,
  loading: false,
  striped: false,
  bordered: false,
  showTableSetting: false,
  useSearchForm: true,
  title: '用户列表',
  titleHelpMessage: "温馨提醒",
  formConfig: {
    autoSubmitOnEnter: true,
    // 关闭折叠
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

  api: async (params) => {
    params.size = params.pageSize
    console.log(params);
    let data = await adminList(params)
    return {items: data.list, total: data.total}
  },
  columns: columns,
  pagination: {pageSize: 10},
  actionColumn: {
    width: 160,
    title: "Action",
    dataIndex: "action"
    // slots: { customRender: 'action' },
  }
});


function handleDelete(record: Recordable) {
  createConfirm({
    iconType: "error",
    title: "提示",
    content: "你正在进行删除操作...",
    onOk: async () => {
      adminDelete({id:record.id}).then(res => {
        success("删除成功");
        reload();
      });
    }
  });
}

</script>

<style scoped lang="less">

</style>
