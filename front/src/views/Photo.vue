<template>
  <div>
    <Buttons />
    <Gallery />
    <Photo />
  </div>
</template>

<script>
import Buttons from '@/components/Buttons.vue';
import Gallery from '@/components/Gallery.vue';
import Photo from '@/components/Photo.vue';

export default {
  name: 'photo',
  components: {
    Buttons,
    Gallery,
    Photo,
  },
};
</script>