<template>
  <div class="container mt-2">
    <nav class="navbar navbar-expand-lg navbar-light bg-light rounded">
      <div class="container-fluid ">
        <a class="navbar-brand" href="#">Home</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item me-2">
              <button type="button" class="btn btn-outline-success" @click="showInput()"> Add </button>
            </li>
            <li class="nav-item me-2 ">
              <router-link :to="`/photo/random?q=${Math.floor(Math.random()*10000)}`">
                <button type="button" class="btn btn-outline-primary"> Magic! </button>
              </router-link>
            </li>
          </ul>
          <!--<form class="d-flex">
            <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search">
            <button class="btn btn-outline-success" type="submit">Search</button>
          </form>-->
        </div>
      </div>
    </nav>
    <div v-if="isUploading">
      <div class="mb-3">
        <label for="formName" class="form-label" >Photo description</label>
        <textarea class="form-control" id="formName" rows="1" v-model="description"></textarea>
      </div>
      <div class="mb-3">
        <label for="formFile" class="form-label">Only jpgs are allowed :)</label>
        <input class="form-control" type="file" id="formFile" accept="image/*" @change="uploadImage($event)">
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import apiurl from "@/path.js"

const axios_instance = axios.create({
  baseURL: apiurl + "upload",
})

export default {
  name: 'Buttons',
  data() {
    return {
      isUploading: false,
      description: "",
      item:{
          image : null,
          imageUrl: null,
      }
    };
  },

  methods: {
    showInput(){
      this.isUploading = !this.isUploading
    },

    uploadImage(event) {

      let data = new FormData();
      data.append('name', 'my-picture');
      data.append('description', this.description)
      data.append('file', event.target.files[0]); 

      let config = {
        header : {
          'Content-Type' : 'image/jpeg'
        }
      }

      axios_instance.post('', data, config).then(
        (response) => {
          console.log('image upload response > ', response)
        }
      )
  }
  },
};
</script>

<style>
  .panel {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 10px;
      grid-auto-rows: 100px;
  }
  button {
      background-color: #4CAF50;
      color: white;
      height: 3rem;
      border-radius: 0.3rem;
      font-size: 1rem;
  }
  .create{
      grid-column: 1;
  }
  .magic {
      grid-column: 2;
  }
  
</style>