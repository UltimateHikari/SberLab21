<template>
  <div class="panel">
    <button class="create"> Create </button>
    <button class="magic"> Magic! </button>
  </div>
</template>

<style>
  .panel {
      display: grid;
      grid-column-gap: 10rem;
      margin: auto 5rem;
  }
  button {
      background-color: #4CAF50;
      color: white;
      height: 3rem;
      border-radius: 0.5rem;
      max-width: 50rem;
      font-size: 1rem;
  }
  .create{
      grid-column: 1;
  }
  .magic {
      grid-column: 2;
  }
  
</style>