import {
  TableCaption,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from "@/components/ui/table";
import { Table } from "@/components/ui/table";
import { useContracts } from "@/src/hooks/useContracts";
import TruncatedAddress from "../../common/truncated-address";
import { Badge } from "@/components/ui/badge";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";

export default function VerifiedContracts() {
  const { formattedContracts } = useContracts();

  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle>Verified Contracts</CardTitle>
      </CardHeader>
      <CardContent>
        <Table>
          <TableCaption>Verified Contracts</TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead>Contract Name</TableHead>
              <TableHead>Confirmed</TableHead>
              <TableHead>Contract Address</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {formattedContracts.map((contract, i) => (
              <TableRow key={i}>
                <TableCell className="font-medium">{contract.name}</TableCell>
                <TableCell>
                  <Badge
                    variant={contract.confirmed ? "success" : "destructive"}
                  >
                    {contract.confirmed ? "Yes" : "No"}
                  </Badge>
                </TableCell>
                <TableCell>
                  <TruncatedAddress address={contract.address} />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  );
}
