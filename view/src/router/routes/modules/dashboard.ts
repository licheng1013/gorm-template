import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const dashboard: AppRouteModule = {
  path: '/dashboard',
  name: 'Dashboard',
  component: LAYOUT,
  redirect: '/dashboard/analysis',
  meta: {
    hideChildrenInMenu: true,
    orderNo: 10,
    icon: 'ion:grid-outline',
    title: t('routes.dashboard.dashboard'),
  },
  children: [
    {
      path: 'analysis',
      name: 'HomeView',
      component: () => import('/@/views/custom/HomeView.vue'),
      meta: {
        icon: 'ion:grid-outline',
        title: "首页",
        hideMenu: true,
      },
    },
  ],
};

export default dashboard;
