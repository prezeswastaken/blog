<script setup lang="ts">
defineProps<{
  title: string;
  body: string;
}>();
</script>

<template>
  <div class="p-5 bg-emerald-500">
    <div class="flex flex-col">
      <p class="text-lg">{{ title }}</p>
      <div class="font-mono font-normal">{{ body }}</div>
    </div>
  </div>
</template>
