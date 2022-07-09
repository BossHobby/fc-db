import { createRouter, createWebHistory } from "vue-router";
import TargetList from "../views/TargetListView.vue";
import Target from "../views/TargetView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "targetList",
      component: TargetList,
    },
    {
      path: "/:mgfr/:board",
      name: "target",
      component: Target,
    },
  ],
});

export default router;
