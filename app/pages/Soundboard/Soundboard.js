import React from 'react';
import Dropzone from 'react-dropzone';
import axios from 'axios';

import './Soundboard.scss';

let self;

export default class Soundboard extends React.Component {
    
    constructor() {
        super();
        this.state = {
            uploaded: false,
            percentCompleted: 0,
            password: "",
        }
        self = this;
    }
    
    componentDidMount() {
        this.config = {
            headers: {
                'Content-Type': 'multipart/form-data',
            },
            onUploadProgress: (progressEvent) => {
                this.setState({
                    percentCompleted: Math.round( (progressEvent.loaded * 100) / progressEvent.total ),
                });
            }
        };
    }
    
    onDrop(acceptedFiles, rejectedFiles) {
      if (acceptedFiles.length > 0) {
          self.uploadFile(acceptedFiles[0]);
      }
    }
    
    uploadFile(file) {
        let formData = new FormData();
        formData.append("name", file.name);
        formData.append("file", file);
        formData.append("password", this.state.password);
        
        axios.put("/upload", formData, this.config)
            .then((response) => {
                this.setState({
                    password: "",
                    percentCompleted: 0,
                    uploaded: true,
                    uploadError: undefined,
                });
            }).catch((err) => {
                this.setState({
                    password: "",
                    percentCompleted: 0,
                    uploaded: false,
                    uploadError: "Upload error.",
                });
            });
    }
    
    passwordOnChange(event) {
        this.setState({
            password: event.target.value,
        });
    }
    
    render() {
        return (
            <div className="Soundboard">
            
            <input className="Soundboard__input"
                    type="password"
                    placeholder="Password"
                    value={this.state.password}
                    onChange={this.passwordOnChange.bind(this)}></input>
            <div>
                <Dropzone className="Dropzone"
                        activeClassName="Dropzone--active"
                        onDrop={this.onDrop}
                        multiple={false}
                        accept={"audio/*"}>
                        
                    <div>Drop file here to upload.</div>
                    {this.state.percentCompleted > 0 ? <div>Progress: {this.state.percentCompleted}</div> : ""}
                    {this.state.uploaded ? <div>File uploded!</div> : ""}
                    {this.state.uploadError ? <div>{this.state.uploadError}</div> : ""}
                </Dropzone>
            </div>
            </div>
        )
    }
}