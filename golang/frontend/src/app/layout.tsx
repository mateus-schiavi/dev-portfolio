import React, { ReactNode } from "react";
import Link from "next/link";

export const metadata = {
    title: "Cardápio Digital",
};

type RootLayoutProps = {
    children: ReactNode;
};

export default function RootLayout({ children }: RootLayoutProps) {
    return (
        <html lang="pt-BR">
            <body>
                <nav style={{ padding: "1rem", borderBottom: "1px solid #ccc" }}>
                    <Link href="/" style={{ marginRight: 20 }}>
                        Cardápio
                    </Link>
                    <Link href="/categories" style={{ marginRight: 20 }}>
                        Categorias
                    </Link>
                    <Link href="/ratings" style={{ marginRight: 20 }}>
                        Avaliações
                    </Link>
                </nav>
                <main style={{ padding: "1rem" }}>{children}</main>
            </body>
        </html>
    );
}
