<script setup lang="ts">
import DefaultLayout from "./DefaultLayout.vue";
import Post from "./Post.vue";

type Post = {
  title: string;
  body: string;
};

const post1: Post = { title: "Fuck you", body: "FIshafiohfsoisfhajoihjoi" };
const post2: Post = { title: "Again?", body: "Pizdofofofofofomofosojfs" };
const post3: Post = { title: "FKfff?", body: "Pizdofofofofofomofosojfs" };
const post4: Post = { title: "FKfff?", body: "Pizdofofofofofomofosojfs" };
const posts: Array<Post> = [post1, post2, post3, post4];
</script>

<template>
  <DefaultLayout>
    <div class="flex flex-col gap-5">
      <p>Hello From child!</p>

      <div v-for="(post, index) in posts" :key="index">
        <Post :title="post.title" :body="post.body" />
      </div>
    </div>
  </DefaultLayout>
</template>
