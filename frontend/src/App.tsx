import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import React, { useEffect } from "react";

import Body from "./components/Body";
import Navbar from "./components/Navbar";
import Appoint from "./components/Appoint";
import Login from "./components/Login";

export default function App() {
  const [token, setToken] = React.useState<String>("");

  useEffect(() => {
    const getToken = localStorage.getItem("uid");
    if (getToken) {
      setToken(getToken);
    }
  }, []);
  console.log("Token",token)

  if (!token) {
    return <Login />
  }

  return (

    <Router>
      {token && (
        <>
        <div>

          <Navbar />

          <Switch>

            <Route exact path="/" component={Body} />

            <Route exact path="/list" component={Appoint} />

          </Switch>

        </div>
        </>
      )}
    </Router>

  );

}