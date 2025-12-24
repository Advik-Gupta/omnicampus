"use client";

import * as React from "react";
import Link from "next/link";
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
} from "@/components/ui/navigation-menu";
import { cn } from "@/lib/utils";

const StudentNavbar = () => {
  return (
    <nav className="w-full border-b bg-background">
      <div className="mx-auto flex h-14 max-w-7xl items-center justify-between px-6">
        <Link href="/dashboard" className="text-lg font-semibold">
          OmniCampus
        </Link>

        <NavigationMenu>
          <NavigationMenuList className="gap-2">
            <NavItem href="/dashboard">Dashboard</NavItem>

            <NavigationMenuItem>
              <NavigationMenuTrigger>Academics</NavigationMenuTrigger>
              <NavigationMenuContent>
                <ul className="grid w-[260px] gap-2 p-3">
                  <DropdownItem
                    href="/courses"
                    title="Courses"
                    description="View registered courses"
                  />
                  <DropdownItem
                    href="/timetable"
                    title="Timetable"
                    description="Weekly class schedule"
                  />
                  <DropdownItem
                    href="/assignments"
                    title="Assignments"
                    description="Submissions & deadlines"
                  />
                  <DropdownItem
                    href="/grades"
                    title="Grades"
                    description="Marks & performance"
                  />
                </ul>
              </NavigationMenuContent>
            </NavigationMenuItem>

            <NavItem href="/events">Events</NavItem>

            <NavigationMenuItem>
              <NavigationMenuTrigger>More</NavigationMenuTrigger>
              <NavigationMenuContent>
                <ul className="grid w-[240px] gap-2 p-3">
                  <DropdownItem
                    href="/announcements"
                    title="Announcements"
                    description="College & faculty notices"
                  />
                  <DropdownItem
                    href="/calendar"
                    title="Academic Calendar"
                    description="Holidays & exams"
                  />
                  <DropdownItem
                    href="/support"
                    title="Support"
                    description="Help & grievances"
                  />
                </ul>
              </NavigationMenuContent>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>

        <Link
          href="/profile"
          className="rounded-full border px-4 py-1.5 text-sm hover:bg-accent"
        >
          Profile
        </Link>
      </div>
    </nav>
  );
};

export default StudentNavbar;

/* ------------------ Helpers ------------------ */

const NavItem = ({
  href,
  children,
}: {
  href: string;
  children: React.ReactNode;
}) => (
  <NavigationMenuItem>
    <Link href={href} legacyBehavior passHref>
      <NavigationMenuLink
        className={cn(
          "rounded-md px-3 py-2 text-sm font-medium hover:bg-accent"
        )}
      >
        {children}
      </NavigationMenuLink>
    </Link>
  </NavigationMenuItem>
);

const DropdownItem = ({
  href,
  title,
  description,
}: {
  href: string;
  title: string;
  description: string;
}) => (
  <li>
    <Link href={href} className="block rounded-md p-2 hover:bg-accent">
      <div className="text-sm font-medium">{title}</div>
      <p className="text-xs text-muted-foreground">{description}</p>
    </Link>
  </li>
);
