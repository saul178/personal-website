import axios from "axios";

const axiosInstance = axios.create({
  // TODO: replace this with a proper domain name
  baseURL: "/api/",
})

export default axiosInstance
