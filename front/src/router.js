import { createRouter, createWebHistory } from "vue-router"
import Home from "@/views/Home.vue"

const routeInfos = [
    {
        path : "/",
        name: "Home",
        component : Home,
    },
]

const router = createRouter({
    history : createWebHistory(),
    routes : routeInfos
})

export default router;