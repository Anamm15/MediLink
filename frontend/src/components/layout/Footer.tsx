import { ShieldCheck } from "lucide-react";
import Link from "next/link";

export const Footer = () => {
  return (
    <footer className="bg-gray-800 text-white">
      <div className="container mx-auto px-4 py-12">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
          {/* Kolom 1: Logo & About */}
          <div className="col-span-2 md:col-span-1">
            <Link
              href="/"
              className="flex items-center gap-2 text-xl font-semibold"
            >
              <ShieldCheck className="h-7 w-7 text-cyan-400" />
              HealthApp
            </Link>
            <p className="mt-4 text-sm text-gray-400">
              Solusi kesehatan digital terpercaya untuk Anda dan keluarga.
            </p>
          </div>
          {/* Kolom lainnya */}
          <div>
            <h3 className="font-semibold">Layanan</h3>
            <ul className="mt-4 space-y-2 text-sm text-gray-400">
              <li>
                <Link href="/doctors" className="hover:text-white">
                  Cari Dokter
                </Link>
              </li>
              <li>
                <Link href="/pharmacy" className="hover:text-white">
                  Apotek Online
                </Link>
              </li>
              <li>
                <Link href="/articles" className="hover:text-white">
                  Artikel Kesehatan
                </Link>
              </li>
            </ul>
          </div>
          <div>
            <h3 className="font-semibold">Tentang</h3>
            <ul className="mt-4 space-y-2 text-sm text-gray-400">
              <li>
                <Link href="/about" className="hover:text-white">
                  Tentang Kami
                </Link>
              </li>
              <li>
                <Link href="/contact" className="hover:text-white">
                  Kontak
                </Link>
              </li>
              <li>
                <Link href="/privacy" className="hover:text-white">
                  Kebijakan Privasi
                </Link>
              </li>
            </ul>
          </div>
          {/* Tambahkan kolom lain jika perlu */}
        </div>
        <div className="mt-12 border-t border-gray-700 pt-8 text-center text-sm text-gray-500">
          © {new Date().getFullYear()} HealthApp. All rights reserved.
        </div>
      </div>
    </footer>
  );
};
