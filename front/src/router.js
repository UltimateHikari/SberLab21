import { createRouter, createWebHistory } from "vue-router"
import Home from "@/views/Home.vue"
import Photo from "@/views/Photo.vue"

const routeInfos = [
    
    {
        path : "/",
        name: "Home",
        component : Home,
    },
    {
        path : "/photo/:id",
        name: "Photo",
        component: Photo,
        props: route => ({ query: route.query.q })
    },
]

const router = createRouter({
    history : createWebHistory(),
    routes : routeInfos
})

export default router;