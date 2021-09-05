<template>
  <h1>GoVue Todo List</h1>
  <div class="todoList" v-for="{ id, title, createdAt } in todos" :key="id">
    <p>{{ title }}</p>
    <p>{{ createdAt }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import axios from "axios";

type Todo = {
  id: string;
  title: string;
  createdAt: string;
  updatedAt: string;
};

export default defineComponent({
  name: "App",
  data() {
    return {
      todos: [] as Todo[],
    };
  },
  async mounted() {
    const { data } = await axios.get("/todo/all");
    const formattedData = data.map(
      ({
        ID,
        Title,
        CreatedAt,
        UpdatedAt,
      }: {
        ID: string;
        Title: string;
        CreatedAt: string;
        UpdatedAt: string;
      }) => ({
        id: ID,
        title: Title,
        createdAt: CreatedAt,
        updatedAt: UpdatedAt,
      })
    ) as Todo[];
    this.todos = formattedData;
  },
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
