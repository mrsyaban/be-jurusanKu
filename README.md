<h2 align="center">
  Server-side JurusanKu<br/>
</h2>
<hr>

> part of Rakuten Team (Hackfest 2024)

<p align="center"> 
  <img src="https://img.shields.io/badge/-Go-333333?style=flat-square&logo=Go">
  <img src="https://img.shields.io/badge/-Docker-333333?style=flat-square&logo=docker"/>
  <img src="https://img.shields.io/badge/-PostgreSQL-333333?style=flat-square&logo=postgresql">
</p>

## Table of Contents
- [Table of Contents](#table-of-contents)
- [General Information](#general-information)
- [Member List](#member-list)
- [Technologies Used](#technologies-used)
- [Setup](#setup)
- [Usage](#usage)
- [Documentation](#documentation)
- [Project Status](#project-status)
- [](#)

<a name="general-information"></a>

## General Information
The system to be developed is a simulation and recommendation system for Indonesian students, specifically designed to be disability-friendly, with the aim of assisting individuals with disabilities in choosing their academic majors. The system will offer courses to simulate lectures related to specific majors. It will be capable of displaying top recommendations based on the user's personality, interests, and talents. Additionally, the system will provide an open discussion platform to connect students seeking information about a particular major with current students or professors in that field.  

<a name="member-list"></a>

## Member List

| Nama                  | Email                       |
| --------------------- | --------------------------- |
| Fajar Maulana Herawan | 13521080@std.stei.itb.ac.id |
| Gevyndo Gunawan       | 18221076@std.stei.itb.ac.id |
| Ilmagita Nariswari    | 18221101@std.stei.itb.ac.id |
| Muhammad Rizky Sya'ban| 13521119@std.stei.itb.ac.id |

<a name="technologies-used"></a>

## Technologies Used
- gin - version 1.9.1
- viper - version 1.18.2
- crypto - version 0.18.0
- pq - version 1.10.9
- jwt-go - version 3.2.0
- docker - version 4.22.1

<a name="setup"></a>

## Setup
You can setup your project by cloning this repository and install the libraries above.

```bash
go get
```

<a name="usage"></a>

## Usage
You can run the server by using the command below.

```bash
make db
```

```bash
make createdb
```

```bash
make server
```

<a name="documentation"> </a>
## Documentation
[postman collection](docs\jurusanKu.postman_collection.json)

<a name="project-status"> </a>

## Project Status
Project is: _on development_

<a name="contact"></a>

## 
<h4 align="center">
  Created by @Rakuten<br/>
  2024
</h4>
<hr>
