<template>
  <h1>{{ target.manufacturer }}-{{ target.board }}</h1>
</template>

<script lang="ts">
import { useTargetStore } from "@/stores/target";
import { mapActions, mapState } from "pinia";
import { defineComponent } from "vue";

export default defineComponent({
  computed: {
    ...mapState(useTargetStore, ["target"]),
  },
  methods: {
    ...mapActions(useTargetStore, ["fetchTarget"]),
  },
  created() {
    this.fetchTarget(
      this.$route.params.mgfr as string,
      this.$route.params.board as string
    );
  },
});
</script>

<style></style>
