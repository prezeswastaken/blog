<script setup lang="ts">
import { onMounted, ref } from "vue";
import DefaultLayout from "./DefaultLayout.vue";
import Post from "./Post.vue";
import CreatingToggle from "./CreatingToggle.vue";

type Post = {
  Title: string;
  Body: string;
  Id: number;
};

const postsArray = ref<Post[]>([]);

let creating = ref(false);

function toggleCreating() {
  creating.value = !creating.value;
  console.log(creating.value);
}

onMounted(() => {
  fetch("//localhost:3000/api/get-all")
    .then((response) => response.json())
    .then((data) => (postsArray.value = data))
    .catch((err) => console.log(err.message));
});
</script>

<template>
  <DefaultLayout>
    <CreatingToggle @toggle-creating="toggleCreating" />
    <div class="flex flex-col gap-5">
      <p>Hello From child!</p>

      <div v-for="post in postsArray" :key="post.Id">
        <Post :title="post.Title" :body="post.Body" />
      </div>
    </div>
  </DefaultLayout>
</template>
