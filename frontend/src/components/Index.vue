<script setup lang="ts">
import { onMounted } from "vue";
import DefaultLayout from "./DefaultLayout.vue";
import Post from "./Post.vue";

type Post = {
  title: string;
  body: string;
};

let postsArray: any;

onMounted(() => {
  // Use the native fetch API to send a GET request
  console.log("starting fetching");
  fetch("//localhost:3000/api/get-all")
    .then((response) => {
      // Check if the response is ok
      if (response.ok) {
        // Return the response as JSON
        console.log("ok");
        return response.json();
      } else {
        // Throw an error with the status text
        throw new Error(response.statusText);
      }
    })
    .then((data) => {
      // Set the data state with the JSON data
      data.value = data;
      console.log(data[0]["Title"]);
      postsArray = data;

      // Parse the JSON string into an array of objects
      //const posts = JSON.parse(data);

      // Cast each object as a Post type
      //postsArray = posts.map((post: Post) => post as Post);
      console.log(postsArray);
      // Set the loading state to false
    })
    .catch((error) => {
      // Set the error state with the error message
      error.value = error.message;
      // Set the loading state to false
    });
});
</script>

<template>
  <DefaultLayout>
    <div class="flex flex-col gap-5">
      <p>Hello From child!</p>

      <div v-for="post in postsArray" :key="post.id">
        <Post :title="post.Title" :body="post.Body" />
      </div>
    </div>
  </DefaultLayout>
</template>
