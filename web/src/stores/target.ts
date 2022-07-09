import { Repo, type Target } from "./api/repo";
import { defineStore } from "pinia";

export const useTargetStore = defineStore({
  id: "target",
  state: () => ({
    target: {} as Target,
  }),
  actions: {
    fetchTarget(mgfr: string, board: string) {
      return Repo.fetchTarget(mgfr, board).then((t) => (this.target = t));
    },
  },
});
