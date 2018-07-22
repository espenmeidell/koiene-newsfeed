import React, { Component } from "react"
import { Fabric } from "office-ui-fabric-react/lib/Fabric"
import Editor from "./components/editor/Editor"

class App extends Component {
  render() {
    return (
      <Fabric>
        <div className="App" style={style}>
          <Editor />
        </div>
      </Fabric>
    )
  }
}

const style = {
  display: "flex",
  justifyContent: "center",
}

export default App
