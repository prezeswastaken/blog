<script setup lang="ts">
defineProps<{
  message: string;
}>();
</script>

<template>
  <div class="fixed top-0 left-1/2 -translate-x-1/2 translate-y-6">
    <button
      class="p-2 text-emerald-100 bg-emerald-900 rounded-3xl border-emerald-100 duration-300 hover:text-emerald-500"
      @click="$emit('toggle-creating')"
    >
      {{ message }}
    </button>
  </div>
</template>
