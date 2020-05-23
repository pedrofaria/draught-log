import PublicLayout from './views/public/layout.svelte'
import IndexView from './views/public/index.svelte'

const routes = [
    {
        name: '/client',
        component: PublicLayout,
        nestedRoutes: [
            { name: 'index', component: IndexView },
        ]
    },
]

export { routes }