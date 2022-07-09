import { Repo, type TargetListEntry } from "./api/repo";
import { defineStore } from "pinia";

export const useTargetListStore = defineStore({
  id: "targetList",
  state: () => ({
    targets: [] as TargetListEntry[],
  }),
  actions: {
    fetchTargetList() {
      if (this.targets.length == 0) {
        return Repo.fetchTargetList().then((l) => (this.targets = l));
      }
    },
  },
});
