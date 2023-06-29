import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';

const dashboard: AppRouteModule = {
  path: '/admin',
  name: 'AdminView',
  component: LAYOUT,
  redirect: '/admin',
  meta: {
    hideChildrenInMenu: true,
    orderNo: 11,
    icon: 'mdi:user-outline',
    title: "用户",
  },
  children: [
    {
      path: '',
      name: 'AdminView',
      component: () => import('/@/views/custom/AdminView.vue'),
      meta: {
        title: "用户",
        hideMenu: true,
      },
    },
  ],
};

export default dashboard;
