<template>
  <div id="app">
    <h1>GoVue Todo List</h1>
    <form method="POST" class="todo-form" v-on:submit.prevent="createTodo">
      <input
        v-model="todo"
        type="text"
        name="todo"
        class="todo-input"
        placeholder="what to do?"
      />
    </form>
    <div method="POST" class="todos" v-for="{ id, title } in todos" :key="id">
      <div class="todo">
        <form v-if="id === editedTodoId" v-on:submit.prevent="updateTodo">
          <input type="text" name="editedTodoTitle" v-model="editedTodoTitle" />
        </form>
        <p v-else>{{ title }}</p>
        <button @click="cancel" v-if="id === editedTodoId">Cancel</button>
        <button @click="updateTodo" v-if="id === editedTodoId">Confirm</button>
        <button @click="editTodo(id, title)" v-else>Edit</button>
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

.todo-form {
  width: 100%;
}

.todo-input {
  width: 100%;
}

.todos {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 512px;
}

.todo {
  width: 100%;
  display: flex;
}
</style>
