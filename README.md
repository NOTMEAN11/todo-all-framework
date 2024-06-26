# Todo App all frameworks

this is a todo app with all frameworks that i know and i will learn in the future

## Database Table

```sql
id int PRIMARY KEY NOT NULL,
title text NOT NULL,
completed boolean NOT NULL DEFAULT false
```

- id is a number that is auto incremented
- title is the title of the todo
- completed is a boolean that is false by default

## Routes

- `GET /todos` - returns all todos
- `GET /todos/:id` - returns a todo by id
- `POST /todos` - creates a todo
- `PATCH /todos/:id` - updates a todo by id
- `DELETE /todos/:id` - deletes a todo by id

## Frameworks

#### Frontend

- [React](https://react.dev/) - Here is the [repo](https://github.com/NOTMEAN11/todo-all-framework/tree/master/frontend/react-todo) for the react app
  <!-- - [Next](https://nextjs.org/)
  - Here is the [repo]() for the pages route
  - Here is the [repo]() for the app route -->
- [Vue](https://vuejs.org/) - Here is the [repo](https://github.com/NOTMEAN11/todo-all-framework/tree/master/frontend/vue-todo) for the vue app
  <!-- - [Nuxt](https://nuxtjs.org/) - Here is the [repo]() for the nuxt app -->
- [Svelte](https://svelte.dev/) - Here is the [repo](https://github.com/NOTMEAN11/todo-all-framework/tree/master/frontend/svelte-todo) for the svelte app
  <!-- - [SvelteKit](https://kit.svelte.dev/) - Here is the [repo]() for the sveltekit app -->

#### Backend

- [Node HTTP](https://nodejs.org/) - Here is the [repo](https://github.com/NOTMEAN11/todo-all-framework/tree/master/backend/node-http-todo)
- [Express](https://expressjs.com/) - Here is the [repo](https://github.com/NOTMEAN11/todo-all-framework/tree/master/backend/express-todo) for the express app
- [Fastify](https://www.fastify.io/) - Here is the [repo](https://github.com/NOTMEAN11/todo-all-framework/tree/master/backend/fastify-todo) for the fastify app
- [Nest](https://nestjs.com/) - Here is the [repo](https://github.com/NOTMEAN11/todo-all-framework/tree/master/backend/nest-todo) for the nest app
- [Gin](https://gin-gonic.com/) - Here is the [repo](https://github.com/NOTMEAN11/todo-all-framework/tree/master/backend/gin-todo) for the gin app
- [Fiber](https://gofiber.io/) - Here is the [repo](https://github.com/NOTMEAN11/todo-all-framework/tree/master/backend/fiber-todo) for the fiber app
<!-- - [Echo](https://echo.labstack.com/) - Here is the [repo]() for the echo app
- [Axum](https://docs.rs/axum/latest/axum/) - Here is the [repo]() for the axum app -->

<!-- #### Fullstack
- [Next](https://nextjs.org/)
  - Here is the [repo]() for the pages route
  - Here is the [repo]() for the app route
- [Nuxt](https://nuxtjs.org/) - Here is the [repo]() for the nuxt app
- [SvelteKit](https://kit.svelte.dev/) - Here is the [repo]() for the sveltekit app
- [Django](https://www.djangoproject.com/) - Here is the [repo]() for the django app -->
