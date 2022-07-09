<template>
  <table class="table">
    <thead>
      <tr>
        <th>MCU</th>
        <th>Manufacturer</th>
        <th>Board</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="t of targets"
        :key="t.board"
        @click="navigate(t)"
        class="board-row"
      >
        <td>{{ t.mcu }}</td>
        <td>{{ t.manufacturer }}</td>
        <td>{{ t.board }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import type { TargetListEntry } from "@/stores/api/repo";
import { useTargetListStore } from "@/stores/targetList";
import { mapActions, mapState } from "pinia";
import { defineComponent } from "vue";

export default defineComponent({
  computed: {
    ...mapState(useTargetListStore, ["targets"]),
  },
  methods: {
    ...mapActions(useTargetListStore, ["fetchTargetList"]),
    navigate(t: TargetListEntry) {
      this.$router.push(`/${t.manufacturer}/${t.board}`);
    },
  },
  created() {
    this.fetchTargetList();
  },
});
</script>

<style>
.board-row {
  cursor: pointer;
}
</style>
