package com.server.javaserver;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Hour {
    private Integer hoursToAdd;

    public Integer getHoursToAdd() {
        return hoursToAdd;
    }

    public void setHoursToAdd(Integer hoursToAdd) {
        this.hoursToAdd = hoursToAdd;
    }
}