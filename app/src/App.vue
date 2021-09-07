<template>
  <div id="app" class="w-full max-w-lg">
    <h1 class="text-3xl font-bold">GoVue Todo List</h1>
    <form method="POST" class="w-full my-6" v-on:submit.prevent="createTodo">
      <input
        v-model="todo"
        type="text"
        name="todo"
        class="
          w-full
          border-2
          px-3
          py-2
          border-gray-300
          rounded-lg
          outline-none
          focus:ring focus:ring-gray-200
        "
        placeholder="what to do?"
      />
    </form>
    <div
      method="POST"
      class="w-full flex flex-col"
      v-for="{ id, title } in todos"
      :key="id"
    >
      <div
        class="
          w-full
          flex flex-row
          items-center
          justify-between
          border-2 border-gray-300
          px-3
          py-4
          my-2
          rounded-xl
          space-x-4
        "
      >
        <form class="w-full flex" v-if="id === editedTodoId" v-on:submit.prevent="updateTodo">
          <input
            type="text"
            name="editedTodoTitle"
            v-model="editedTodoTitle"
            class="
              w-full
              border-2
              px-3
              py-2
              border-gray-300
              rounded-lg
              outline-none
              focus:ring focus:ring-gray-200
            "
          />
        </form>
        <p v-else>{{ title }}</p>
        <div class="flex flex-shrink-0 space-x-2">
          <button
            class="text-gray-800 py-1 px-2 font-medium"
            @click="cancel"
            v-if="id === editedTodoId"
          >
            Cancel
          </button>
          <button
            class="
              border-2 border-gray-800
              text-gray-800
              bg-gray-50
              py-1
              px-2
              rounded-md
              font-medium
            "
            @click="updateTodo"
            v-if="id === editedTodoId"
          >
            Confirm
          </button>
          <button
            class="
              bg-gray-800
              border-2 border-gray-800
              text-gray-50
              py-1
              px-2
              font-medium
              rounded-md
            "
            @click="editTodo(id, title)"
            v-else
          >
            Edit
          </button>
        </div>
      </div>
    </div>
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
      todo: "",
      editedTodoId: "",
      editedTodoTitle: "",
    };
  },
  mounted() {
    this.fetchTodos();
  },
  methods: {
    async fetchTodos() {
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
    async createTodo() {
      try {
        const payload = {
          title: this.todo,
        };
        const { data } = await axios.post("/todo/create", payload);
        const formattedData = {
          id: data?.ID,
          title: data?.Title,
          createdAt: data?.CreatedAt,
          updatedAt: data?.UpdatedAt,
        } as Todo;
        this.todos.push(formattedData);
        this.todo = "";
      } catch (err) {
        console.log(err);
      }
    },
    editTodo(id: string, title: string) {
      this.editedTodoId = id;
      this.editedTodoTitle = title;
    },
    cancel() {
      this.editedTodoId = "";
      this.editedTodoTitle = "";
    },
    async updateTodo() {
      try {
        const payload = {
          id: this.editedTodoId,
          title: this.editedTodoTitle,
        };
        console.log(payload);
        const { data } = await axios.put(`/todo/update/${payload.id}`, payload);
        const formattedData = {
          id: data?.ID,
          title: data?.Title,
          createdAt: data?.CreatedAt,
          updatedAt: data?.UpdatedAt,
        } as Todo;
        const updatedTodosIndex = this.todos.findIndex(
          (todo) => todo.id === formattedData.id
        );
        this.todos[updatedTodosIndex] = formattedData;
        console.log(this.todos);
        this.cancel();
      } catch (error) {
        console.log(error);
      }
    },
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
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
