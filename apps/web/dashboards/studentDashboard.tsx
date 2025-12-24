"use client";

import React from "react";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  CardDescription,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import { ScrollArea } from "@/components/ui/scroll-area";

const CourseCard = ({ name, code }: { name: string; code: string }) => (
  <Card className="hover:shadow-md transition">
    <CardContent className="p-4 space-y-2">
      <h3 className="font-semibold">{name}</h3>
      <p className="text-sm text-muted-foreground">{code}</p>
      <Button variant="outline" size="sm">
        View
      </Button>
    </CardContent>
  </Card>
);

const Announcement = ({ title, tag }: { title: string; tag: string }) => (
  <div className="flex items-center justify-between">
    <p className="text-sm">{title}</p>
    <Badge variant="secondary">{tag}</Badge>
  </div>
);

const UpcomingItem = ({ title, date }: { title: string; date: string }) => (
  <div className="flex justify-between text-sm">
    <span>{title}</span>
    <span className="text-muted-foreground">{date}</span>
  </div>
);

const StudentDashboard = () => {
  return (
    <div className="p-6 space-y-6">
      {/* Main Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Left Column */}
        <div className="lg:col-span-2 space-y-6">
          {/* Courses */}
          <Card>
            <CardHeader>
              <CardTitle>Your Courses</CardTitle>
              <CardDescription>Currently enrolled subjects</CardDescription>
            </CardHeader>
            <CardContent className="grid sm:grid-cols-2 gap-4">
              <CourseCard name="Data Structures" code="CSE201" />
              <CourseCard name="Operating Systems" code="CSE301" />
              <CourseCard name="Mathematics III" code="MAT202" />
              <CourseCard name="Computer Networks" code="CSE303" />
              <CourseCard name="Computer Networks" code="CSE303" />
              <CourseCard name="Computer Networks" code="CSE303" />
            </CardContent>
          </Card>

          {/* Announcements */}
          <Card>
            <CardHeader>
              <CardTitle>Announcements</CardTitle>
              <CardDescription>From faculty & administration</CardDescription>
            </CardHeader>
            <CardContent>
              <ScrollArea className="h-[180px] space-y-4">
                <Announcement
                  title="Mid-Sem Exams Schedule Released"
                  tag="Exam"
                />
                <Announcement
                  title="Assignment Deadline Extended"
                  tag="Academic"
                />
                <Announcement title="Holiday on Friday" tag="Notice" />
              </ScrollArea>
            </CardContent>
          </Card>
        </div>

        {/* Right Column */}
        <div className="space-y-6">
          {/* Calendar */}
          <Card>
            <CardHeader>
              <CardTitle>Calendar</CardTitle>
            </CardHeader>
            <CardContent>
              <Calendar mode="single" />
            </CardContent>
          </Card>

          {/* Upcoming */}
          <Card>
            <CardHeader>
              <CardTitle>Upcoming</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
              <UpcomingItem title="OS Quiz" date="Tomorrow" />
              <UpcomingItem title="Math Assignment" date="In 3 days" />
              <UpcomingItem title="CN Test" date="Next Week" />
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default StudentDashboard;
