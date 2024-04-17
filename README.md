<a name="readme-top"></a>
<!--*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/z0057393/QRCodeGenerator">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/23/Go_Logo_Aqua.svg/1200px-Go_Logo_Aqua.svg.png" alt="Logo" width="180" height="80">
  </a>

  <h3 align="center">API Proxy Rotation</h3>

  <p align="center">
    API Rest in GO for QRCode Generation 
    <br />
    <a href="https://github.com/z0057393/QRCodeGenerator"><strong>Explore the docs »</strong></a>
  </p>
</div>

<!-- GETTING STARTED -->
## Getting Started


### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Docker : [link](https://www.docker.com/)
  ```sh
    docker --version                                                          
    Docker version 24.0.7, build afdd53b
  ```

* Visual studio code : [link](https://code.visualstudio.com/)
* Dev container extension : [link](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
  


### Development


You could make any modification in the code by using dev container extension 

Make sure your docker engine is running.

Dev container should detect the .devcontainer folder and ask you to reopen the container. 

If not use the VS Code Command Palette to reopen it 

You can access the VS Code Command Palette in a number of ways. 
Shift + Command + P (Mac) / Ctrl + Shift + P (Windows/Linux). 



<!-- USAGE EXAMPLES -->
## Usage

Start the API 

```sh
    go run main.go
```

In localhost the API should be accesible with the URL : localhost:8080
Use the route (POST method): /api/generate 
with the parameters size and content 



<!-- ROADMAP -->
## Roadmap

- [ ] Add Changelog
- [ ] Add License
- [ ] Securing the API 

<!-- - 
- [x] Add back to top links
- [ ] Add Additional Templates w/ Examples
- [ ] Add "components" document to easily copy & paste sections of the readme
- [ ] Multi-language Support
    - [ ] Chinese
    - [ ] Spanish

See the [open issues](https://github.com/othneildrew/Best-README-Template/issues) for a full list of proposed features (and known issues). -->




<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request





<!-- LICENSE -->
## License

See `LICENSE.txt` for more information.


[contributors-shield]: https://img.shields.io/github/contributors/z0057393/QRCodeGenerator.svg?style=for-the-badge
[contributors-url]: https://github.com/z0057393/QRCodeGenerator/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/z0057393/QRCodeGenerator.svg?style=for-the-badge
[forks-url]: https://github.com/z0057393/QRCodeGenerator/network/members
[stars-shield]: https://img.shields.io/github/stars/z0057393/QRCodeGenerator.svg?style=for-the-badge
[stars-url]: https://github.com/z0057393/QRCodeGenerator/stargazers
[issues-shield]: https://img.shields.io/github/issues/z0057393/QRCodeGenerator.svg?style=for-the-badge
[issues-url]: https://github.com/z0057393/QRCodeGenerator/issues
[license-shield]: https://img.shields.io/github/license/z0057393/QRCodeGenerator.svg?style=for-the-badge
[license-url]: https://github.com/z0057393/QRCodeGenerator/blob/master/LICENSE.txt