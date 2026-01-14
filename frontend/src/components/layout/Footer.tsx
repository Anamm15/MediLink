import { ShieldCheck } from "lucide-react";
import Link from "next/link";

export const Footer = () => {
  return (
    <footer className="bg-gray-800 text-white">
      <div className="container mx-auto px-4 py-12">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
          <div className="col-span-2 md:col-span-1">
            <Link
              href="/"
              className="flex items-center gap-2 text-xl font-semibold"
            >
              <ShieldCheck className="h-7 w-7 text-cyan-400" />
              HealthApp
            </Link>
            <p className="mt-4 text-sm text-gray-400">
              Trusted digital healthcare solution for you and your family.
            </p>
          </div>
          <div>
            <h3 className="font-semibold">Services</h3>
            <ul className="mt-4 space-y-2 text-sm text-gray-400">
              <li>
                <Link href="/doctors" className="hover:text-white">
                  Find a Doctor
                </Link>
              </li>
              <li>
                <Link href="/pharmacy" className="hover:text-white">
                  Online Pharmacy
                </Link>
              </li>
              <li>
                <Link href="/articles" className="hover:text-white">
                  Health Articles
                </Link>
              </li>
            </ul>
          </div>
          <div>
            <h3 className="font-semibold">About</h3>
            <ul className="mt-4 space-y-2 text-sm text-gray-400">
              <li>
                <Link href="/about" className="hover:text-white">
                  About Us
                </Link>
              </li>
              <li>
                <Link href="/contact" className="hover:text-white">
                  Contact
                </Link>
              </li>
              <li>
                <Link href="/privacy" className="hover:text-white">
                  Privacy Policy
                </Link>
              </li>
            </ul>
          </div>
        </div>
        <div className="mt-12 border-t border-gray-700 pt-8 text-center text-sm text-gray-500">
          © {new Date().getFullYear()} HealthApp. All rights reserved.
        </div>
      </div>
    </footer>
  );
};
