import React, { Component } from "react"
import { Fabric } from "office-ui-fabric-react/lib/Fabric"
import { BrowserRouter as Router, Route, Link } from "react-router-dom"
import { DefaultButton } from "office-ui-fabric-react"

import Editor from "./components/Editor/Editor"
import Home from "./components/Home/Home"

class App extends Component {
  render() {
    return (
      <Fabric>
        <Router>
          <div className="App" style={style}>
            <nav style={navStyle}>
              <Link to="/">Home</Link> <Link to="/editor">Editor</Link>
            </nav>
            <main style={mainStyle}>
              <Route exact path="/" component={Home} />
              <Route path="/editor" component={Editor} />
            </main>
          </div>
        </Router>
      </Fabric>
    )
  }
}

const style = {
  display: "flex",
  flexDirection: "column",
}

const navStyle = {
  padding: "0.5em",
}

const mainStyle = {
  display: "flex",
  justifyContent: "center",
}

export default App
