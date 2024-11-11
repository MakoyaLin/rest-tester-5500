package com.server.javaserver;

import com.fasterxml.jackson.annotation.JsonProperty;

public class User {
    private Integer id;
    private String name;
    private int hoursWorked;

    public User(@JsonProperty("id") Integer id,
                @JsonProperty("name") String name,
                @JsonProperty("hoursWorked") Integer hoursWorked) {
        this.id = id != null ? id : 0;
        this.name = name != null ? name : "";
        this.hoursWorked = hoursWorked != null ? hoursWorked : 0;
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getHoursWorked() {
        return hoursWorked;
    }

    public void setHoursWorked(int hoursWorked) {
        this.hoursWorked = hoursWorked;
    }
}