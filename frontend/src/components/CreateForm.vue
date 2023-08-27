<script setup lang="ts">
const emit = defineEmits(["toggle-creating"]);
let title = "";
let body = "";

function handleSubmit() {
  console.log(title);
  console.log(body);

  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      title: title,
      body: body,
    }),
  };
  fetch("//localhost:3000/api/create", requestOptions).then(() => {
    emit("toggle-creating");
  });
}
</script>

<template>
  <form @submit.prevent="handleSubmit">
    <div
      class="flex fixed top-0 left-0 justify-center items-center w-screen h-screen bg-black bg-opacity-50"
    >
      <div
        class="flex flex-col gap-5 p-10 w-1/2 h-3/4 font-semibold bg-opacity-90 rounded-lg bg-emerald-950"
      >
        <input
          v-model="title"
          class="p-5 placeholder-emerald-300 bg-emerald-500 rounded-lg"
          placeholder="Your title goes here"
        />
        <textarea
          v-model="body"
          placeholder="Write your pasta here"
          class="p-5 h-full placeholder-emerald-300 bg-emerald-500 rounded-lg"
        ></textarea>
        <div class="flex justify-center">
          <button
            class="p-2 mt-5 text-emerald-100 bg-emerald-900 rounded-lg duration-300 hover:text-emerald-300"
          >
            Create
          </button>
        </div>
      </div>
    </div>
  </form>
</template>
