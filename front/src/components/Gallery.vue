<template>
  <div class="container mt-2">
    <div class="gallery">
      <div class="gallery-panel"
          v-for="photo in photos"
          :key="photo.id">
        <router-link :to="`/photo/${photo.id}`">
          <img :src="thumbUrl(photo.id)">
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import apiurl from "@/path.js"

const axios_instance = axios.create({
  baseURL: apiurl + "list",
})

export default {
  name: 'Gallery',
  data() {
    return {
      photos: [],
    };
  },
  created() {
    console.log("get to " + apiurl + "list")
    axios_instance.get().then((result) => {
      this.photos = result.data;
    }, error => {
      console.error(error);
    });
  },

  methods: {
    thumbUrl(id) {
      var location = apiurl + id + ".jpg";
      console.log("pulling photo from " + location);
      return location;
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
    //margin: 2rem auto;
    //padding: 0 5rem;
  }  
  .gallery-panel img {
    width: 100%;
    height: 22vw;
    object-fit: cover;
    border-radius: 0.75rem;
  }
</style>