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
        path : "/photo/:filename",
        name: "Photo",
        component: Photo,
    }
]

const router = createRouter({
    history : createWebHistory(),
    routes : routeInfos
})

export default router;