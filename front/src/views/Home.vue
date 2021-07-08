<template>
  <h1> Home </h1>
  <Buttons />
  <Gallery />
</template>

<script>
import Gallery from '@/components/Gallery.vue';
import Buttons from '@/components/Buttons.vue';

export default {
  name: 'home',
  components: {
    Buttons,
    Gallery,
  },
};
</script>

<style>
    h1{
        text-align: center;
    }
</style>