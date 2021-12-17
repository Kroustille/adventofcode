package challenge

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Challenge struct {
	all_bits          string
	current_bit_index int
	version_sum       int
}

const (
	SUM_OPERATOR_TYPE_ID          = 0
	PRODUCT_OPERATOR_TYPE_ID      = 1
	MINIMUM_OPERATOR_TYPE_ID      = 2
	MAXIMUM_OPERATOR_TYPE_ID      = 3
	VALUE_TYPE_ID                 = 4
	GREATER_THAN_OPERATOR_TYPE_ID = 5
	LESS_THAN_OPERATOR_TYPE_ID    = 6
	EQUAL_OPERATOR_TYPE_ID        = 7
	FLAG_LAST_VALUE               = "0"
	LENGTH_TYPE_ID_TOTAL_LENGTH   = "0"
)

func (c Challenge) ConvertBitsToDecimal(bits string) int {
	decimal_value, err := strconv.ParseInt(bits, 2, len(bits)+1)
	if err != nil {
		log.Fatal(err)
	}
	return int(decimal_value)
}

func (c Challenge) BuildBitString(line string) string {
	splitted_line := strings.Split(line, "")
	final_bits := ""
	for _, bit := range splitted_line {
		bit_value, err := strconv.ParseUint(bit, 16, 4)
		if err != nil {
			log.Fatal(err)
		}

		bit_string := fmt.Sprintf("%04b", bit_value)
		final_bits += bit_string
	}

	return final_bits
}

func (c *Challenge) Read(n int) string {
	value := c.all_bits[c.current_bit_index : c.current_bit_index+n]
	c.current_bit_index += n
	return value
}

func (c *Challenge) ReadVersion() int {
	// log.Println("read :version")
	version_bits := c.Read(3)
	// log.Println("value:version =", version_bits)
	return c.ConvertBitsToDecimal("0" + version_bits)
}

func (c *Challenge) ReadType() int {
	// log.Println("read :type")
	type_bits := c.Read(3)
	// log.Println("value:type =", type_bits)
	return c.ConvertBitsToDecimal("0" + type_bits)
}

func (c *Challenge) DecodePayloadLength() int {
	// log.Println("read :payload_length")
	payload_length := c.Read(15)
	// log.Println("value:payload_length:", payload_length)
	return c.ConvertBitsToDecimal(payload_length)
}

func (c *Challenge) DecodePacketsCount() int {
	// log.Println("read :packets_count")
	packets_count := c.Read(11)
	// log.Println("value:packets_count:", packets_count)
	return c.ConvertBitsToDecimal(packets_count)
}

func (c *Challenge) DecodePacketsCountOperator(n int) []int {
	values := []int{}
	for i := 0; i < n; i++ {
		values = append(values, c.Decode())
	}

	return values
}

func (c *Challenge) DecodePayloadLengthOperator(n int) []int {
	bytes_read := 0
	saved_bit_index := c.current_bit_index
	values := []int{}
	for bytes_read < n {
		values = append(values, c.Decode())
		bytes_read = c.current_bit_index - saved_bit_index
	}

	return values
}

func (c Challenge) ComputeSum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum
}

func (c Challenge) ComputeProduct(values []int) int {
	product := 1
	for _, value := range values {
		product *= value
	}

	return product
}

func (c Challenge) ComputeMinimum(values []int) int {
	minimum := math.MaxInt
	for _, value := range values {
		if value < minimum {
			minimum = value
		}
	}

	return minimum
}

func (c Challenge) ComputeMaximum(values []int) int {
	maximum := 0
	for _, value := range values {
		if value > maximum {
			maximum = value
		}
	}

	return maximum
}

func (c Challenge) ComputeGreaterThan(values []int) int {
	if values[0] > values[1] {
		return 1
	}

	return 0
}

func (c Challenge) ComputeLessThan(values []int) int {
	if values[0] < values[1] {
		return 1
	}

	return 0
}

func (c Challenge) ComputeEqual(values []int) int {
	if values[0] == values[1] {
		return 1
	}

	return 0
}
func (c *Challenge) DecodeOperatorPacket(type_id int) int {
	length_type_id := c.Read(1)
	// log.Println("length_type_id =", length_type_id)
	var values []int
	if length_type_id == LENGTH_TYPE_ID_TOTAL_LENGTH {
		payload_length := c.DecodePayloadLength()
		// log.Println("payload_length =", payload_length)
		values = c.DecodePayloadLengthOperator(payload_length)
	} else {
		packets_count := c.DecodePacketsCount()
		// log.Println("packets_count =", packets_count)
		values = c.DecodePacketsCountOperator(packets_count)
	}

	switch type_id {
	case SUM_OPERATOR_TYPE_ID:
		return c.ComputeSum(values)
	case PRODUCT_OPERATOR_TYPE_ID:
		return c.ComputeProduct(values)
	case MINIMUM_OPERATOR_TYPE_ID:
		return c.ComputeMinimum(values)
	case MAXIMUM_OPERATOR_TYPE_ID:
		return c.ComputeMaximum(values)
	case GREATER_THAN_OPERATOR_TYPE_ID:
		return c.ComputeGreaterThan(values)
	case LESS_THAN_OPERATOR_TYPE_ID:
		return c.ComputeLessThan(values)
	case EQUAL_OPERATOR_TYPE_ID:
		return c.ComputeEqual(values)
	}

	return 0
}

func (c *Challenge) DecodeValuePacket() int {
	// log.Println("DecodeValuePacket")
	is_end_of_packet := false
	bits_string := ""
	bits_read_count := 6 // version and type length that are already read

	for !is_end_of_packet {
		// log.Println("read :flag")
		flag_bit := c.Read(1)
		bits_read_count += 1

		// log.Println("value:flag =", flag_bit)
		if flag_bit == FLAG_LAST_VALUE {
			is_end_of_packet = true
		}

		var current_bits_string string
		// log.Println("read :value")
		current_bits_string = c.Read(4)
		bits_read_count += 4
		// log.Println("value:value =", current_bits_string)
		bits_string += current_bits_string
	}

	// log.Println("value:bits_string =", c.ConvertBitsToDecimal(bits_string))

	return c.ConvertBitsToDecimal(bits_string)
}

func (c *Challenge) Decode() int {
	// log.Println("Decode")

	version := c.ReadVersion()
	// log.Println("version =", version)
	c.version_sum += version
	type_id := c.ReadType()
	// log.Println("type_id =", type_id)

	packet_value := 0
	switch type_id {
	case VALUE_TYPE_ID:
		packet_value = c.DecodeValuePacket()
		break
	default:
		packet_value = c.DecodeOperatorPacket(type_id)
	}

	return packet_value
}
