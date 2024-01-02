import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/IndexPage.vue') }],
  },
  {
    path: '/login',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/auth/LoginPage.vue') },
    ],
  },
  {
    path: '/register',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/auth/RegisterPage.vue') },
    ],
  },
  {
    path: '/profile',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/auth/ProfilePage.vue') },
    ],
  },

  {
    path: '/new_collection',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        component: () => import('pages/collection/NewCollection.vue'),
      },
    ],
  },
  {
    path: '/collection/:collectionId',
    props: true,
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        props: true,
        component: () => import('pages/collection/CollectionView.vue'),
      },
    ],
  },
  {
    path: '/collection/:collectionId/dashboard',
    props: true,
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        props: true,
        component: () => import('pages/collection/CollectionView.vue'),
      },
    ],
  },
  {
    path: '/collection/:collectionId/specimens',
    props: true,
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        props: true,
        component: () => import('pages/collection/SpecimenList.vue'),
      },
    ],
  },
  {
    path: '/collection/:collectionId/new/specimen',
    props: true,
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        props: true,

        component: () => import('pages/collection/NewSpecimen.vue'),
      },
    ],
  },
  {
    path: '/collection/:collectionId/bulk_import',
    props: true,
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        props: true,
        component: () => import('pages/collection/BulkImport.vue'),
      },
    ],
  },
  {
    path: '/collection/:collectionId/specimen/:specimenId',
    props: true,
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        props: true,

        component: () => import('pages/collection/SpecimenView.vue'),
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
