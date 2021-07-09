<template>
  <div class="gallery">
    <div class="gallery-panel"
         v-for="photo in photos"
         :key="photo.id">
      <router-link :to="`/photo/${photo.filename}`">
        <img :src="thumbUrl(photo.filename)">
      </router-link>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import ip from "@/ip.json";


const axios_instance = axios.create({
  baseURL: "http://" + ip.ip + "/photos/list",
})

export default {
  name: 'Gallery',
  data() {
    return {
      photos: [],
    };
  },
  created() {
    axios_instance.get().then((result) => {
      this.photos = result.data;
    }, error => {
      console.error(error);
    });
  },

  methods: {
    thumbUrl(filename) {
      return require(`../assets/images/thumbnails/${filename}`);
    },
  },
};
</script>

<style>
  .gallery {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(20rem, 1fr));
    grid-gap: 1rem;
    max-width: 80rem;
    margin: 2rem auto;
    padding: 0 5rem;
  }  
  .gallery-panel img {
    width: 100%;
    height: 22vw;
    object-fit: cover;
    border-radius: 0.75rem;
  }
</style>