## XJCO3211_Distributed_Systems

![golang](https://img.shields.io/badge/Go-00ADD8.svg?style=for-the-badge&logo=Go&logoColor=white)

This README provides an overview of the implementation of three tasks aimed at creating a distributed system based on serverless computing using Microsoft Azure.

### Module: Distributed Systems (XJCO3211)
This project entails building a distributed system with serverless functions on Azure to emulate an IoT network for collecting and processing environmental data. The project is organized into three primary components: data simulation, statistical analysis, and configuring a realistic scenario with automated processes.

### Task

#### Task 1: Simulated Data Collection

Gather environmental data from 20 sensors via an HTTP-triggered Azure Function. The collected data is saved in an Azure SQL database, making it available for subsequent analysis.

#### Task 2: Data Statistics

Deploy an Azure Function with an SQL trigger to detect updates in the Azure SQL database holding sensor data. This function calculates the minimum, maximum, and average values for each sensor, providing insights into any unusual environmental conditions.

#### Task 3: Realistic Scenario Implementation

Set up automated data collection and analysis by combining a time trigger and SQL trigger within Azure Functions. This configuration enables new data entries to automatically initiate statistical analysis.

### Azure Setup

- Azure SQL Server: erfeiyu
- Azure SQL Database: erfeiyu
- Azure Functions: erfeiyu
    - SimulatedDataTask1 -- HTTP Trigger
    - SimulatedData -- Timer Trigger
    - Statistics -- SQL Trigger
- Application Insights: erfeiyu202411030642