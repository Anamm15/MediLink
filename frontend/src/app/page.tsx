import { Navbar } from "@/components/layout/Navbar";
import { HeroSection } from "./components/HeroSection";
import { HowItWorksSection } from "./components/HowItWorksSection";
import { TestimonialSection } from "./components/TestimonialSection";
import { Footer } from "@/components/layout/Footer";
import { StatsSection } from "./components/StatsSection";
import { FeaturesSection } from "./components/FeaturesSection";
import { CtaSection } from "./components/CTASection";
// Import komponen lain seperti FeaturesSection dan CtaSection jika sudah dibuat

export default function LandingPage() {
  return (
    <>
      <Navbar />
      <HeroSection />
      <StatsSection />
      <HowItWorksSection />
      <FeaturesSection />
      <TestimonialSection />
      <CtaSection />
      <Footer />
    </>
  );
}
