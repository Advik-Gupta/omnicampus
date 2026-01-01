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
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { cn } from "@/lib/utils";
import { useRouter } from "next/navigation";
import { useAuthStore } from "@/store/auth.store";
import { useEffect } from "react";

const StudentNavbar = () => {
  const router = useRouter();
  const user = useAuthStore((state) => state.user);

  useEffect(() => {
    console.log("User data in Navbar:", user);
  }, [user]);

  const handleLogout = () => {
    localStorage.removeItem("omnicampus-auth");
    cookieStore.delete("auth_token");
    router.push("/login");
  };

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

        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Avatar className="cursor-pointer h-9 w-9">
              <AvatarImage src="https://i.pravatar.cc/150?img=12" />
              <AvatarFallback>
                {user?.name
                  ?.split(" ")
                  .map((n) => n[0])
                  .join("")
                  .slice(0, 2)
                  .toUpperCase()}
              </AvatarFallback>
            </Avatar>
          </DropdownMenuTrigger>

          <DropdownMenuContent
            align="end"
            className="w-56 rounded-xl p-0 overflow-hidden"
          >
            <div className="px-4 py-3 border-b bg-muted/40">
              <p className="text-sm font-semibold leading-tight">
                {user?.name}
              </p>
              <p className="text-xs text-muted-foreground">
                Reg No: {user?.register_number}
              </p>
            </div>

            <div className="p-1">
              <DropdownMenuItem asChild>
                <Link
                  href="/profile"
                  className="flex items-center gap-2 px-3 py-2 rounded-md"
                >
                  Profile
                </Link>
              </DropdownMenuItem>

              <DropdownMenuItem
                className="flex items-center gap-2 px-3 py-2 rounded-md text-red-600 focus:text-red-600"
                onClick={handleLogout}
              >
                Logout
              </DropdownMenuItem>
            </div>
          </DropdownMenuContent>
        </DropdownMenu>
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
