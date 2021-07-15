<template>
  <div class="lightbox" @click.self="close">
    <div v-if="this.loadingStatus" class='loading-div'>
      <div class="spinner-border text-info m-5" role="status">
        <span class="sr-only">Loading...</span>
      </div>
    </div>
    <div v-else class='loading-div'>
      <img class="rounded" :src="imgsrc">      
    </div>
    <div class="lightbox-info">
      <div class="lightbox-info-inner">
        Info
      </div>
    </div>
  </div>
</template>

<script>
import apiurl from "@/path.js"

export default {
  name: 'Photo',
  props: {
    query: {
      type: Number,
      default: 1000
    }
  },
  data() {
    return {
      imgsrc: "",
      loadingStatus: true,
    };
  },
  computed: {
    photo() {
      return this.photos.find((photo) => {
        return photo.id === Number(this.$route.params.id);
      });
    },
  },
  created() {
    console.log(this.loadingStatus)

    var myImage = new Image();
    myImage.src = this.photoUrl(this.$route.params.id);
    myImage.onload = () => {
      console.log("loaded!")
      this.imgsrc = myImage.src
      this.loadingStatus = false;
    }
    
  },
  methods: {
    photoUrl(id) {
      console.log(this.$route.query.q + " ; " + id);
      console.log("query: " + this.query + "|");
      var location = apiurl + id + ".jpg" + "?q=" + this.$route.query.q;
      console.log("pulling photo from " + location);
      return location;
    },
    close(){
        this.$router.push("/")
    }
  },
};
</script>

<style>
  .lightbox {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.8);
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-gap: 2rem;
  }  
  .loading-div {
    margin: auto;
    width: 100%;
    grid-column-start: 2;
  } 
  .lightbox-info {
    margin: auto 2rem auto 0;
  }  
  .lightbox-info-inner {
    background-color: #FFFFFF;
    display: inline-block;
    padding: 2rem;
  }
</style>