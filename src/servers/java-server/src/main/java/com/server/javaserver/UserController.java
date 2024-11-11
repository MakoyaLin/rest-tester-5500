package com.server.javaserver;

import org.springframework.util.StringUtils;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.atomic.AtomicInteger;

@RestController
@RequestMapping("/users")
public class UserController {

    private final List<User> users = new ArrayList<>();
    private final AtomicInteger nextId = new AtomicInteger(1);

    @GetMapping
    public List<User> getAllUsers() {
        System.out.println("GET /users - Fetching all users");
        return users;
    }

    @GetMapping("/{id}")
    public Optional<User> getUserById(@PathVariable Integer id) {
        System.out.println("GET /users/" + id + " - Fetching user with ID: " + id);
        Optional<User> first = users.stream().filter(u -> u.getId().equals(id)).findFirst();
        if (!first.isPresent()) {
            throw new ResourceNotFoundException("User not found");
        }
        return first;
    }

    @DeleteMapping
    public List<User> deleteAllUsers() {
        System.out.println("DELETE /users - Deleting all users");
        users.clear();
        nextId.set(1);
        return users;
    }

    @PostMapping
    public User addUser(@RequestBody User newUser) {
        System.out.println("POST /users - Adding new user");
        if (StringUtils.isEmpty(newUser.getName())) {
            throw new IllegalArgumentException("Name is required and must be a non-empty string");
        }
        newUser.setId(nextId.getAndIncrement());
        users.add(newUser);
        return newUser;
    }

    @PutMapping("/{id}")
    public User updateUserById(@PathVariable Integer id, @RequestBody User updatedUser) {
        System.out.println("PUT /users/" + id + " - Updating user with ID: " + id);
        User user = getUserById(id).orElseThrow(() -> new ResourceNotFoundException("User not found"));
        if (!StringUtils.isEmpty(updatedUser.getName())) {
            user.setName(updatedUser.getName());
        }
        return user;
    }

    @PatchMapping("/{id}")
    public User updateHoursWorked(@PathVariable Integer id, @RequestBody Hour hour) {
        Integer hoursToAdd = hour.getHoursToAdd();
        System.out.println("PATCH /users/" + id + "/hours - Updating hours for user with ID: " + id);
        User user = getUserById(id).orElseThrow(() -> new ResourceNotFoundException("User not found"));
        user.setHoursWorked(user.getHoursWorked() + hoursToAdd);
        return user;
    }

    @DeleteMapping("/{id}")
    public User deleteUserById(@PathVariable Integer id) {
        System.out.println("DELETE /users/" + id + " - Deleting user with ID: " + id);
        Optional<User> first = users.stream().filter(u -> u.getId().equals(id)).findFirst();
        if (first.isPresent()) {
            users.removeIf(u -> u.getId().equals(id));
            return first.get();
        } else {
            throw new ResourceNotFoundException("User not found");
        }
    }
}