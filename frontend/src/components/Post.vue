<script setup lang="ts">
let props = defineProps({
  title: String,
  body: String,
  id: Number,
});
const emit = defineEmits(["fetch"]);

function handleClick() {
  console.log("My id is " + props.id);
  handleDelete();
}

function handleDelete() {
  console.log("Deleting " + props.id + " post.");

  const requestOptions = {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Id: props.id,
    }),
  };
  fetch("//localhost:3000/api/delete", requestOptions).then(() => {
    emit("fetch");
  });
}
</script>

<template>
  <div class="p-5 bg-emerald-500 rounded-xl">
    <div class="flex flex-col">
      <div class="flex justify-between">
        <p class="text-lg">{{ title }}</p>
        <p
          class="text-xl duration-300 hover:text-red-500 hover:cursor-grab"
          @click="handleClick"
        >
          󰆴
        </p>
      </div>
      <div class="font-mono font-normal">{{ body }}</div>
    </div>
  </div>
</template>
