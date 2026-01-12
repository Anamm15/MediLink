"use client";
import { Lock } from "lucide-react";
import Script from "next/script";
import { useBooking } from "../hooks/useBooking";
import { CreateBookingRequest } from "@/types/appointment.type";
import { getSnapTokenFromUrl } from "@/helpers/midtrans";
import { toast } from "sonner";

declare global {
  interface Window {
    snap: any;
  }
}

interface PaymentFlowProps {
  consultationFee: number;
  payload: CreateBookingRequest;
}

export const PaymentFlow = ({ consultationFee, payload }: PaymentFlowProps) => {
  const adminFee = 2500;
  const total = consultationFee + adminFee;
  const { mutateAsync: createBooking } = useBooking();

  const handlePayment = async () => {
    // LANGKAH 1: Minta token dari backend Anda
    try {
      const response = await createBooking(payload);
      const token = getSnapTokenFromUrl(response.payment_url);
      if (window.snap && token) {
        window.snap.pay(token, {
          onSuccess: function (result: any) {
            console.log("success", result);
            toast.message("Payment success", {
              duration: 3000,
            });
          },
          onPending: function (result: any) {
            console.log("pending", result);
            toast.message("Payment failed, please try again", {
              duration: 3000,
            });
          },
          onError: function (result: any) {
            console.log("error", result);
            toast.message("Payment failed, please try again", {
              duration: 3000,
            });
          },
          onClose: function () {
            console.log(
              "customer closed the popup without finishing the payment"
            );
            toast.message("Payment failed, please try again", {
              duration: 3000,
            });
          },
        });
      }
    } catch (error) {}
  };

  return (
    <>
      <Script
        src="https://app.sandbox.midtrans.com/snap/snap.js"
        data-client-key={process.env.NEXT_PUBLIC_MIDTRANS_CLIENT_KEY}
        strategy="afterInteractive"
      />
      <div className="bg-white p-6 rounded-xl border border-gray-200 shadow-sm flex flex-col h-full">
        <div>
          <h2 className="text-xl font-bold text-gray-800 mb-4">Bill Details</h2>
          <div className="space-y-3 text-sm">
            <div className="flex justify-between">
              <p className="text-gray-500">Consultation Fee</p>
              <p className="font-medium text-gray-800">
                Rp {consultationFee.toLocaleString("id-ID")}
              </p>
            </div>
            <div className="flex justify-between">
              <p className="text-gray-500">Service Fee</p>
              <p className="font-medium text-gray-800">
                Rp {adminFee.toLocaleString("id-ID")}
              </p>
            </div>
            <div className="border-t border-dashed my-3"></div>
            <div className="flex justify-between text-base">
              <p className="font-semibold text-gray-800">Total</p>
              <p className="font-bold text-cyan-600 text-lg">
                Rp {total.toLocaleString("id-ID")}
              </p>
            </div>
          </div>
        </div>

        <div className="mt-auto pt-8">
          <button
            onClick={handlePayment}
            className="w-full bg-slate-800 text-white font-bold py-3.5 rounded-lg text-lg hover:bg-slate-700 transition-colors duration-200 flex items-center justify-center gap-2"
          >
            <Lock className="w-5 h-5" />
            Pay Now
          </button>
          <p className="text-xs text-gray-400 text-center mt-3">
            You will be redirected to a secure payment page.
          </p>
        </div>
      </div>
    </>
  );
};
