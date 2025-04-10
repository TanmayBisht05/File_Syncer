import {create} from "zustand";
import axiosInstance from "../lib/axios";
import toast from "react-hot-toast";

export const useAuthStore=create((set)=>({
    authUser: null,
    isLoggingIn:false,
    isSigningUp:false,
    isCheckingAuth: true,


    checkAuth: async () => {
      try {
        const res = await axiosInstance.get("/auth/check");
        set({ authUser: res.data });
      } catch (error) {
        set({ authUser: null });
      } finally {
        set({ isCheckingAuth: false });
      }
    },
  



    login: async(data)=>{
        set({isLoggingIn: true});

        try {
            const res=await axiosInstance.post("/auth/login",data);
            set({authUser: res.data});
            toast.success("Logged in Successfully");
            
        } catch (error) {
            toast.error(error.response?.data?.error || "Login Failed"); 
        }

        finally{
            set({isLoggingIn: false});
        }
    },


    signup: async (data) => {
        set({ isSigningUp: true });
        try {
          console.log("testing1");
          const res = await axiosInstance.post("/auth/signup", data);
          set({ authUser: res.data });
          toast.success("Signup successful");
        } catch (error) {
          toast.error(error.response?.data?.error || "Signup failed");
        } finally {
          set({ isSigningUp: false });
        }
    },
    
    logout: async () => {
        try {
          await axiosInstance.post("/auth/logout");
          set({ authUser: null });
          toast.success("Logged out successfully");
        } catch (error) {
          toast.error(error.response?.data?.error || "Logout failed");
        }
    },
}))