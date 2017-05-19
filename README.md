# Replicated QA Automation Coding Challenge

## Purpose

This coding challenge will involve testing two unauthenticated endpoints on GitHub's public API that accept and return JSON.  You will be using the Ginkgo Testing framework with the Gomega assertion library which is what we use here at Replicated  Additionally, I've set up this challenge using other tools we use such as Docker and Makefiles.  Please take your time and add up to 5 scenarios per endpoint.  We realize there's potentially an infinite amount of edge case scenarios, so what we're really looking for is what you think would be the best scenarios to test.

## Details for submission

Please fork and clone this GitHub repository, then submit a pull request back into it with your changes.  Feel free to modify or add whatever you need including addition of other packages, your own libraries, etc.  Please replace this README.md with a list and explanation of your test scenarios.  

## Documentation

[GitHub API Documentation](https://developer.github.com/v3/)

[Ginkgo Documentation](https://onsi.github.io/ginkgo/)

[Gomega Documentation](http://onsi.github.io/gomega/)

[Installing Docker](https://docs.docker.com/)

## Setup for development

Build the Docker container:

```shell
make docker
```

Run the tests:

```shell
make test
```

While the container is running, you will be able to edit the tests on your host system and they will be immediately available after saving inside the Docker container.

## What we review

* **Correctness:** Do your tests compile and run?  Do they actually make calls to the API and verify results?
* **Code Quality:** Is your code maintainable, easy to understand, and conducive for future reuse or expansion?
* **Quality Focused:** Does your choice of test case scenarios provide adequate, given the limits, confidence in the quality of the GitHub API endpoints tested?
* **Communication:** Are you able to defend why you picked your test case scenarios and explain what they do in your README and through other forms of communication?

![Good Luck!](http://i.imgur.com/DHxjAeQ.jpg)
