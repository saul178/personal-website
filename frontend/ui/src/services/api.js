import axios from "axios";

const axiosInstance = axios.create({
  // TODO: replace this with a proper domain name
  baseURL: "http://localhost:8080/api/"
})

export default axiosInstance
