<template>
    <h1> Home </h1>
  <Gallery />
</template>

<script>
import Gallery from '@/components/Gallery.vue';

export default {
  name: 'home',
  components: {
    Gallery,
  },
};
</script>

<style>
    h1{
        text-align: center;
    }
</style>