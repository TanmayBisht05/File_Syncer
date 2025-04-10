import React, {useEffect} from "react";
import {Routes, Route, Navigate} from "react-router-dom";
import {Loader} from "lucide-react";
import {Toaster} from "react-hot-toast";

import "./App.css";

import Navbar from "./components/Navbar";
import LoginPage from "./pages/LoginPage";
import SignUpPage from "./pages/SignUpPage";
import HomePage from "./pages/HomePage";
import ProfilePage from "./pages/ProfilePage";

import {useAuthStore} from "./store/useAuthStore";

const App =()=>{
  const {authUser,checkAuth, isCheckingAuth} =useAuthStore();


  useEffect(()=>{
    checkAuth();
  },[checkAuth]);


  if(isCheckingAuth && !authUser){
    return (
      <div className="centered-loader">
        <Loader/>
      </div>
    );
  }

  return (
    <div>
      <Navbar />
      <div className="page-content">
        <Routes>
          <Route path="/" element={authUser ? <HomePage /> : <Navigate to="/login" />} />
          <Route path="/signup" element={!authUser ? <SignUpPage /> : <Navigate to="/" />} />
          <Route path="/login" element={!authUser ? <LoginPage /> : <Navigate to="/" />} />
          <Route path="/profile" element={<ProfilePage />} />

        </Routes>
      </div>
      <Toaster />
    </div>
  );
  
};



export default App;
